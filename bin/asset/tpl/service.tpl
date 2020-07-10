package {{packageName}}.{{moduleName}}.service;

import com.baomidou.mybatisplus.service.IService;
import {{mainPath}}.common.utils.PageUtils;
import {{packageName}}.{{moduleName}}.entity.{{ className }}Entity;

import java.util.List;
import java.util.Map;

/**
  * {{ table.TableComment }}
  *
  * @author {{ authorName }}
  * @email {{ emailAddress }}
  * @date {{ genTime }}
  */
public interface {{ className }}Service extends IService<{{ className }}Entity> {

    PageUtils queryPage(Map<String, Object> params);

    boolean update({{ className }}Entity entity);

    boolean insert({{ className }}Entity entity);

    void deleteBatchIds(List<String> ids);
}


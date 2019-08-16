package {{packageName}}.{{moduleName}}.dao;

import {{packageName}}.{{moduleName}}.entity.{{ className }}Entity;
import com.baomidou.mybatisplus.mapper.BaseMapper;
import com.baomidou.mybatisplus.plugins.Page;
import org.apache.ibatis.annotations.Mapper;

import java.util.List;
import java.util.Map;

/**
  * {{ table.TableComment }}
  *
  * @author {{ authorName }}
  * @email {{ emailAddress }}
  * @date {{ genTime }}
  */
@Mapper
public interface {{ className }}Dao extends BaseMapper<{{ className }}Entity> {

    void deleteBatchIds(List<String> ids);

    List<{{ className }}Entity> page(Page<{{ className }}Entity> page,Map params);

}